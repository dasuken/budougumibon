# 14
* shutdownにerrGroup
* errGroup contextでcontext共有
    * 目的は複数のリクエストに対してクライアントのタイムアウトやシャットアウトに対して正しくプロセスを閉じる事
    * 感覚的には`ListenAndServe`でlistenしてAcceptしてgoroutineをフォークしているイメージだった
        * errGroupによるgoroutine生成は何回？１回な気がしてしまうんだけど、、
    * => run(ctx)のctxに対するcancel処理を検知してサーバーをshutdownするのが主目的
        * かつ`ErrServerShutdown`以外のエラーが発生した場合は細くしcontextにメッセージを送信して欲しい
        
まだこの実装の必要性を身に染みている漢字はしない、、、
* Graceful Shutdownの考え方が分かってない？
    * もし決済処理のトランザクション中にsigtermでサーバーのプロセスがとじたら？
        * shutdownを実行し、サーバー側でerror処理
            * errGroupはそのプロセスの終了をまってくれるって話かな？
        *   
    * https://zenn.dev/pyotarou/articles/87d43169e0abe0
    
* 普通の実装を書いてみる
    * タイムアウト処理に対してシャットダウンコマンドを実行し、非同期に実行されているgoroutineを停止できるか試す？
    
* shutdown()
    * `select <- ctx.Done`なら`ctx.Err`
    * `default`は`time.C`で設定した秒数まつ
        * 引数にTimeoutを渡す事でシャットダウン処理を任意の秒数まつ
    