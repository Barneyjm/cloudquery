package appservice

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Plans() *schema.Table {
	return &schema.Table{
		Name:        "azure_appservice_plans",
		Resolver:    fetchPlans,
		Description: "https://learn.microsoft.com/en-us/rest/api/appservice/app-service-environments/list-app-service-plans?tabs=HTTP#appserviceplan",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_appservice_plans", client.Namespacemicrosoft_web),
		Transform:   transformers.TransformWithStruct(&armappservice.Plan{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{client.SubscriptionID},
	}
}

func fetchPlans(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armappservice.NewPlansClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
