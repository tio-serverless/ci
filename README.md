# ci
Trigger Travis CI Daily

定时触发Travis的构建任务,可以在空闲时刻做CI / CT

### 配置
> myTravis-ci.toml


```toml
[repos]
  branch=[""]
  name=["xx"]
```

+ branch 需要构建的分支名称
+ name travis中的工程名称