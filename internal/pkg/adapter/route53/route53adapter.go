package route53

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/nousefreak/go-ddns/internal/pkg/config"
	log "github.com/sirupsen/logrus"
)

const TYPE = "aws-route53"

type Route53Adapter struct {
}

func (a *Route53Adapter) SetIP(ip string, config *config.ConfigSet) error {
	options := config.GetOptions()
	zoneID := options["zoneId"]
	record := config.GetRecord()

	sess := session.New()
	if _, ok := options["AWS_ACCESS_KEY_ID"]; ok {
		sess = session.Must(session.NewSession(&aws.Config{
			Credentials: credentials.NewStaticCredentials(
				options["AWS_ACCESS_KEY_ID"], options["AWS_SECRET_ACCESS_KEY"], ""),
		}))
	}

	svc := route53.New(sess)
	result, err := svc.ChangeResourceRecordSets(&route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String("UPSERT"),
					ResourceRecordSet: &route53.ResourceRecordSet{
						Name: aws.String(record),
						ResourceRecords: []*route53.ResourceRecord{{
							Value: aws.String(ip),
						}},
						TTL:  aws.Int64(60),
						Type: aws.String("A"),
					},
				},
			},
			Comment: aws.String("DDNS record updated by route53-ddns"),
		},
		HostedZoneId: aws.String(zoneID),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			log.WithField("type", aerr.Code).Error(aerr.Error())
		} else {
			log.Error(err.Error())
		}
		return err
	}

	log.Debug("Waiting for record set change")
	svc.WaitUntilResourceRecordSetsChanged(&route53.GetChangeInput{
		Id: result.ChangeInfo.Id,
	})

	return nil
}
