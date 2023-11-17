#### fugu.secrets 

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | bigint unsigned | PRI | NO | auto_increment |  |
| 2 | created_at |  | datetime(3) |  | YES |  |  |
| 3 | updated_at |  | datetime(3) |  | YES |  |  |
| 4 | deleted_at |  | datetime(3) | MUL | YES |  |  |
| 5 | secret_id |  | bigint |  | YES |  |  |
| 6 | author_id |  | bigint |  | YES |  |  |
| 7 | view_level |  | bigint |  | YES |  |  |
| 8 | timestamp |  | bigint |  | YES |  |  |
| 9 | views |  | bigint |  | YES |  |  |
| 10 | content |  | varchar(255) |  | YES |  |  |
| 11 | images |  | json |  | YES |  |  |
| 12 | status |  | tinyint(1) |  | YES |  |  |
