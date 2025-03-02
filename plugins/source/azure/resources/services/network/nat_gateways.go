package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func NatGateways() *schema.Table {
	return &schema.Table{
		Name:        "azure_network_nat_gateways",
		Resolver:    fetchNatGateways,
		Description: "https://learn.microsoft.com/en-us/rest/api/virtualnetwork/nat-gateways/list?tabs=HTTP#natgateway",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_network_nat_gateways", client.Namespacemicrosoft_network),
		Transform:   transformers.TransformWithStruct(&armnetwork.NatGateway{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{client.SubscriptionID},
	}
}

func fetchNatGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armnetwork.NewNatGatewaysClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListAllPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
