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


    逆向思路： 更新不及时 需要bot去更新好友关系
	1 每次新用户来都获取一下他的好友（粉丝＋关注）（思考下有没有必要存储）
	2 当有非首位用户注册时,梳理关系。判断他和之前所以的站内用户是否有朋友关系。外循环是他的XFriends，内循环是站内用户，
	3 如果有 那就在SiteFriends插入两条记录 一条是此userid为主，匹配到的朋友为副，一条是，朋友为主，他为副。其实不用，插入一条即可，就是搜的时候需要反搜一下：在userid1 搜一下，再在userid2那里搜一下。
	4 秘密发布的时候，检查作者站内朋友，将此秘密id和这些朋友userid加入此表。查的时候，也是查此表，
	5 查userid对应可见的secretsID，查secrets表，将对应id的内容返回给给前端。

-----------
    正向思路： 更新及时 每次都是最新好友关系
	 	如何保证 我关注+取关后的秘密可见更新的及时性？
			用户请求进来的时候（每次登录就只做登录那一次同步，或者，否则每次请求进来都要同步。）
         执行下1，2，获取他的好友，然后与站内用户对比，将此用户在SiteFriends里将朋友全部更新
      （建议采用插入两条的方案，3那里插入两条，即每次以userid1为主去搜，减少开发者心智负担。删除操作也是）。
	错了 这里既然是每次请求都要更新 那么直接以此用户为主键 全删后再插入。即可保证最新。
	然后拿最新的朋友和秘密的作者取交集。responseSecrets =  select * from secrets s where s.author in (select friendid from SiteFriends where basepersonid  =  'userid' )
	将 此 responseSecrets 返回前端。



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


/*
	如何优化 上次获取到数据和这次获取的数据的对比？
	根据数量，数量不一致肯定不同，数量一样的，也有可能有变化。
	测试下hash变没变，因为推特上顺序是不变的。
------------------------------
	最好的办法，搞机器人，监控每一个用户，一旦用户粉丝和关注变了，就更新数据。
------------------------------
利用redis, set 可以很快求差集。
uid001_storedFd,
uid001_freshFd，
将差集找出，查看差集是否在sitefd里边。在的，就增或删，在viewableFd里。
增的了就将所有推文的

-------------------
每一次登陆 一更新的话 ，就可以采用推特的策略。登录的时候 更新一下关系。然后每次翻页都是直接去redis取


---------------------------------------------

搞不来 大概率还是每次请求都需要判断。
那么就只能

刷新策略，
策略1 每次请求都会更新
策略2 只有登录的时候更新一次，登陆：重连钱包。
策略3 前20次是每次请求都会更新，之后用策略2
策略4 只有登录的时候更新一次，之后必须刷新（update friend circle）才会更新。需前端配合。


动作策略 一：
刷新请求进来，根据id获取最新关注和粉丝列表。在redis中与之前的取差集。
SDIFF 旧 新 = 旧的有 新的没有 = 删除的
SDIFF 新 旧 = 新的有 旧的没有 = 新增的
在mysql数据库siteFriends 中将 删除的和新增的状态更新。
然后去秘密表拿秘密 select * from secrets s where s.author in basefriend() order by timestamp desc.

动作策略 二：
刷新请求进来，根据id获取最新关注和粉丝列表。（直接在redis中存最新的，redis都不用存）
然后更新mysql数据库siteFriends


动作策略 三：
siteFriends在redis中，user 也在redis中，
redis的list结构
刷新请求进来，根据id获取最新关注和粉丝列表。与user取交集，更新siteFriend
然后根据siteFriends 中好友 去mysql中找可见秘密











*/
