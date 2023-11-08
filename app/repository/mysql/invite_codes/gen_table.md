#### fugu.invite_codes 

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint unsigned | PRI | NO | auto_increment |  |
| 2 | created_at |  | datetime(3) |  | YES |  |  |
| 3 | updated_at |  | datetime(3) |  | YES |  |  |
| 4 | deleted_at |  | datetime(3) | MUL | YES |  |  |
| 5 | userid |  | bigint |  | YES |  |  |
| 6 | code |  | varchar(255) |  | YES |  |  |
