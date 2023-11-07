#### fugu.secret_images 

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint unsigned | PRI | NO | auto_increment |  |
| 2 | created_at |  | datetime(3) |  | YES |  |  |
| 3 | updated_at |  | datetime(3) |  | YES |  |  |
| 4 | deleted_at |  | datetime(3) | MUL | YES |  |  |
| 5 | secret_id |  | bigint unsigned | MUL | YES |  |  |
| 6 | image_url |  | longtext |  | YES |  |  |
| 7 | secretid |  | bigint unsigned |  | YES |  |  |
| 8 | imageurl |  | varchar(255) |  | YES |  |  |
