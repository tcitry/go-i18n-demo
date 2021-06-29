# 操作步骤

1. 安装`goi18n`命令

```
go get -u github.com/nicksnyder/go-i18n/v2/goi18n
```

2. 生成英文模板 `active.en.toml`

```
goi18n extract
```

3. 根据英文生成中文模板

```
goi18n merge active.*.toml
```

4. 根据中文模板进行翻译，最后进行合并

```
goi18n merge active.*.toml translate.*.toml
```

(踩了一坑：命名为 active.cn.toml 的话翻译无效)