# slack-summarize-reactions

[ダウンロードはこちらから](https://github.com/Kanatani28/slack-summarize-reactions/releases)

Windows: slack-summarize-reactions_windows_x86_64.zip   
Mac: slack-summarize-reactions_darwin_x86_64.tar.gz   
Linux: slack-summarize-reactions_linux_x86_64.tar.gz   

config.ymlとusers.csvに設定をしてから実行ファイルを叩くと  
Slackのメッセージに付与されたリアクションを集計できます。

## Slackの設定
1. [ここ](https://api.slack.com/apps)にアクセスし、Create an Appボタンをクリックする。
2. 任意のApp Nameを入力、対象のworkspaceを選択し、Create Appボタンをクリックする。
![スクリーンショット 2020-04-12 3 20 30](https://user-images.githubusercontent.com/16130443/79051724-c9130300-7c6c-11ea-8a6b-4cb3cc24c527.png)
3. OAuth & PermissionsページでTokenを作成する。
![キャプチャ](https://user-images.githubusercontent.com/16130443/79060170-b32d2e80-7cbc-11ea-8be9-2e0878194a7a.PNG)
4. User Token Scopesを以下のように設定する。
![キャプチャ2](https://user-images.githubusercontent.com/16130443/79060172-b6281f00-7cbc-11ea-91ef-b7495d66d79c.PNG)
5. 適宜ReInstall Appする

## config.yml

|項目|説明|
|--|--|
|token|Slackで設定したトークン|
|target_channel|集計対象チャンネル名|
|search_count|集計対象件数（投稿日が新しいメッセージから取得する）|

## users.csv

Slackのユーザー名。
[こんな感じ](https://github.com/Kanatani28/slack-summarize-reactions/blob/master/users.csv)で一行一人設定する。    
**人によってはスペースある人とかない人とかいるので注意**
