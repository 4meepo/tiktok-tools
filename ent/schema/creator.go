package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Creator holds the schema definition for the Creator entity.
type Creator struct {
	ent.Schema
}

// Fields of the Creator.
func (Creator) Fields() []ent.Field {
	return []ent.Field{
		field.String("xzid").Unique().MaxLen(128).Comment("先知网ID"),
		field.String("unique_id").MaxLen(64).Comment("唯一ID"),
		field.String("nick_name").MaxLen(100).Comment("昵称"),
		field.String("region").MaxLen(16).Comment("地区"),
		field.Uint32("follower_num").Comment("粉丝数"),
		field.String("creator_id").Default("").MaxLen(32).Comment("tiktok creator_id"),
		field.String("creator_oecuid").Default("").MaxLen(32).Comment("tiktok shop creator oecuid"),
		field.JSON("cate1_name_cn", []string{}).Default([]string{}).Comment("先知侧商品类目"),
		field.JSON("tiktok_category", []string{}).Default([]string{}).Comment("tiktok侧商品类目"),
		field.String("email").MaxLen(64).Comment("先知网爬的email"),
		field.String("whatsapp").MaxLen(64).Comment("先知网爬的whatsapp"),
	}
}

// Edges of the Creator.
func (Creator) Edges() []ent.Edge {
	return nil
}

func (Creator) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("xzid").Unique(),
		index.Fields("region", "follower_num"),
	}
}

func (Creator) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
