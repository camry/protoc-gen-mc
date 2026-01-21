# protoc-gen-mc

protoc-gen-mc 是一个 Go 语言的 Protocol Buffers 插件，用于为生成的 PB 对象添加 MongoDB 操作方法。

## 功能

- 为 PB 消息结构体生成 `Set()` 方法，用于收集 MongoDB 的 `$set` 指令字段
- 支持多层级嵌套消息的字段路径生成（如 `user.address.city`）
- 支持 map 类型字段的路径生成（如 `user.tags.keyName` 或 `user.profiles.keyName.fieldName`）
- 生成 MongoDB 更新操作的辅助方法
- 为所有字段生成 `SetXxx()` 方法
- 为数字类型字段生成 `IncrXxx()` 和 `DecrXxx()` 方法
- 为 map 类型字段生成 `SetXxx()`, `InsertXxx()`, `DeleteXxx()`, `ClearXxx()` 方法

## 安装

```bash
go build -o protoc-gen-mc.exe .
```

确保生成的可执行文件在你的 PATH 环境变量中。

## 使用

```bash
protoc --mongo_out=. your_proto_file.proto
```

## 生成的方法示例

对于以下的 proto 定义：

```proto
message User {
    int64 id = 1;
    string name = 2;
    int32 age = 3;
    Address address = 4;
    map<int32, string> tags = 6;
    map<int32, Address> addr = 7;
}
```

插件将生成以下方法：

### Set 方法

- `Set()` - 生成完整的更新对象
- `SetName(v string) bson.M` - 设置单个字段
- `SetAge(v int32) bson.M` - 设置数字字段
- `SetAddress(v *Address) bson.M` - 设置嵌套消息

### 数字字段方法

- `IncrAge(v int32) bson.M` - 增加数值
- `DecrAge(v int32) bson.M` - 减少数值

### Map 字段方法

- `SetTags(v map[int32]string) bson.M` - 替换整个 map
- `InsertTags(k int32, v string) bson.M` - 向 map 插入键值对
- `DeleteTags(k int32) bson.M` - 从 map 删除键
- `ClearTags() bson.M` - 清空整个 map
