const selectFile = () => {
    // FileListオブジェクト取得
    const selectFiles = document.querySelector("#select-file").files

    // Fileオブジェクト取得
    const file = selectFiles[0]

    // FileReaderオブジェクト取得
    const reader = new FileReader()
    reader.readAsText(file)

    // ファイル読み込み完了時の処理
    reader.onload = () => {
      document.querySelector("#output").innerHTML = reader.result
    }

    // ファイル読み込みエラー時の処理
    reader.onerror = () => {
      console.log("ファイル読み込みエラー")
    }
  }