// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TiktokCreatorsColumns holds the columns for the "tiktok_creators" table.
	TiktokCreatorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "creator_id", Type: field.TypeString, Size: 30},
		{Name: "creator_name", Type: field.TypeString, Size: 100},
		{Name: "creator_nickname", Type: field.TypeString, Size: 100},
		{Name: "region", Type: field.TypeString, Size: 8},
		{Name: "product_categories", Type: field.TypeJSON},
		{Name: "follower_count", Type: field.TypeUint32, Default: 0},
		{Name: "video_avg_view_cnt", Type: field.TypeUint32, Default: 0},
		{Name: "video_pub_cnt", Type: field.TypeUint32, Default: 0},
		{Name: "ec_video_avg_view_cnt", Type: field.TypeUint32, Default: 0},
		{Name: "creator_oecuid", Type: field.TypeString, Default: ""},
		{Name: "creator_ttuid", Type: field.TypeString, Default: ""},
	}
	// TiktokCreatorsTable holds the schema information for the "tiktok_creators" table.
	TiktokCreatorsTable = &schema.Table{
		Name:       "tiktok_creators",
		Columns:    TiktokCreatorsColumns,
		PrimaryKey: []*schema.Column{TiktokCreatorsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "tiktokcreator_creator_id",
				Unique:  true,
				Columns: []*schema.Column{TiktokCreatorsColumns[1]},
			},
			{
				Name:    "tiktokcreator_creator_name",
				Unique:  false,
				Columns: []*schema.Column{TiktokCreatorsColumns[2]},
			},
			{
				Name:    "tiktokcreator_region_follower_count",
				Unique:  false,
				Columns: []*schema.Column{TiktokCreatorsColumns[4], TiktokCreatorsColumns[6]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TiktokCreatorsTable,
	}
)

func init() {
}
