# FuguBackend



Backend service of Fugu App.

type single = map[string]string
type storage = map[string]single

/*
基础表：
table Users
table Secrets
关系表：
table XFriends
table SiteFriends
table viewableSecrets 每一个user都有对应的的viewable Secrets.请求的时候 直接从这个表拿。更新时机，每次有秘密发布。

	1 每次新用户来都获取一下他的好友（粉丝＋关注）（思考下有没有必要存储）
	2 当有非首位用户注册时,梳理关系。判断他和之前所以的站内用户是否有朋友关系。外循环是他的XFriends，内循环是站内用户，
	3 如果有 那就在SiteFriends插入两条记录 一条是此userid为主，匹配到的朋友为副，一条是，朋友为主，他为副。其实不用，插入一条即可，就是搜的时候需要反搜一下：在userid1 搜一下，再在userid2那里搜一下。
	4 秘密发布的时候，检查作者站内朋友，将此秘密id和这些朋友userid加入此表。查的时候，也是查此表，
	5 查userid对应可见的secretsID，查secrets表，将对应id的内容返回给给前端。

-----------

	 	如何保证 我关注+取关后的秘密可见更新的及时性？
			用户请求进来的时候（每次登录就只做登录那一次同步，或者，否则每次请求进来都要同步。）执行下1，2，获取他的好友，然后与站内用户对比，将此用户在SiteFriends里将朋友全部更新（建议采用插入两条的方案，3那里插入两条，即每次以userid1为主去搜，减少开发者心智负担。删除操作也是）。
*/
type twitterClient struct {
}
type Xuser struct {
}

func (t *twitterClient) GetXFriends(id string) []Xuser {
	return []Xuser{}
}

var xClient twitterClient

func XFriends(id string) {
	friends := xClient.GetXFriends(id)
	err := repository.save(id, friends)
}
func FgFriends() {

}
