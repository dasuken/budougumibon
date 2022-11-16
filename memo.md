# 14
* shutdownにerrGroup
* errGroup contextでcontext共有
    * 目的は複数のリクエストに対してクライアントのタイムアウトやシャットアウトに対して正しくプロセスを閉じる事
    * 感覚的には`ListenAndServe`でlistenしてAcceptしてgoroutineをフォークしているイメージだった
        * errGroupによるgoroutine生成は何回？１回な気がしてしまうんだけど、、