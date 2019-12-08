# go-clean-architecture-sample

[あえてGo言語でClean Architectureを学ぶ](https://www.aintek.xyz/posts/clean-architecture-lean-from-golang)の写経。

## メモ

記事の通り、以下のようにpackageを分けてみたけど、使い勝手はいいのかもこの写経で体感できると嬉しい。

```zsh
❯ tree .              
.
├── entities
├── frameWorksAndDrivers
├── interfaceAdapter
└── usecases
```

- `user.go`

`ValidatePassword` とかのが好みだけど、名前は変えない

エラーチェックせずにindex accessしてて怖いけど、本質じゃないから省略したのかな。

## References
- [あえてGo言語でClean Architectureを学ぶ](https://www.aintek.xyz/posts/clean-architecture-lean-from-golang)