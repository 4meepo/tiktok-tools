package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// TiktokCreator holds the schema definition for the TiktokCreator entity.
type TiktokCreator struct {
	ent.Schema
}

// Fields of the TiktokCreator.
func (TiktokCreator) Fields() []ent.Field {
	return []ent.Field{
		field.String("creator_id").MaxLen(30),
		field.String("creator_name").MaxLen(100),
		field.String("creator_nickname").MaxLen(100),
		field.String("region").MaxLen(8),
		field.Strings("product_categories").Default([]string{}),
		field.Uint32("follower_count").Default(0),
		field.Uint32("video_avg_view_cnt").Default(0),
		field.Uint32("video_pub_cnt").Default(0),
		field.Uint32("ec_video_avg_view_cnt").Default(0),
		field.String("creator_oecuid").Default(""),
		field.String("creator_ttuid").Default(""),
	}
}

// Edges of the TiktokCreator.
func (TiktokCreator) Edges() []ent.Edge {
	return nil
}

func (TiktokCreator) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("creator_id").Unique(),
		index.Fields("creator_name"),
		index.Fields("region", "follower_count"),
	}
}

func (TiktokCreator) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
