# GAE/Gのgoapp testを便利に使うためのツール

## Version
0.0.1

## 使い方
$ aetesthelper -sdk_path="/Users/yokoyama/golang/go_appengine" apps/sampleproject

$ aetesthelper -sdk_path=`pwd` apps/sampleproject

## 注意事項
1. AppengineのSDKのディレクトリに各プロジェクトがある前提で作っています。
1. プロジェクト以下の全てのディレクトリに対して、goapp testを実行するので、「no buildable Go source files」が大量に出るかもしれません。
1. 引数にはSDKのディレクトリからの**相対パス**でプロジェクトのパスを指定して下さい。
1. 実行した後、**何が起こっても怒らないで下さい**。
1. 著作権はT.Yokoyamaにあります。
