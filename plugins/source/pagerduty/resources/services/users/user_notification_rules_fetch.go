package users

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchUserNotificationRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)
	concreteParent := parent.Item.(pagerduty.User)

	response, err := cqClient.PagerdutyClient.ListUserNotificationRulesWithContext(ctx, concreteParent.ID)
	if err != nil {
		return err
	}

	if len(response.NotificationRules) == 0 {
		return nil
	}

	res <- response.NotificationRules

	return nil
}
