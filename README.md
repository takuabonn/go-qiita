# go-qiita

### 概要

本 CLI アプリケーションは、ユーザーから入力されたタグを使用して Qiita API を検索し、最新から 10 件の記事を取得し、それを csv に記入するものです。

### 機能

- タグの入力

  ユーザーがコマンドラインからタグを入力します。タグは 1 つのみ指定できます。

- Qiita API の検索

  入力されたタグを使用して Qiita API を検索します。API キーは事前に設定しておく必要があります。

- 記事の取得

  検索結果から、最新から 10 件の記事を取得します。

- csv への記入

  取得した記事を csv に記入します
