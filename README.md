# slack-summarize-reactions

[ダウンロードはこちらから](https://github.com/Kanatani28/slack-summarize-reactions/releases)

Windows: slack-summarize-reactions_windows_x86_64.zip   
Mac: slack-summarize-reactions_darwin_x86_64.tar.gz   
Linux: slack-summarize-reactions_linux_x86_64.tar.gz   

config.ymlとusers.csvに設定をしてから実行ファイルを叩くと  
Slackのメッセージに付与されたリアクションを集計できます。

## Slackの設定


## config.yml

|項目|説明|
|--|--|
|token|Slackで設定したトークン|
|target_channel|集計対象チャンネル名|
|search_count|集計対象件数（投稿日が新しいメッセージから取得する）|

## users.csv

Slackのユーザー名。
[こんな感じ](https://github.com/Kanatani28/slack-summarize-reactions/blob/master/users.csv)で設定する。
