# PDF を画像に変換

## 構成

- [gographics/imagick](https://github.com/gographics/imagick)

## 手順

### 1. Imagemagick がインストールされているか確認

- `sudo dpkg -l`
- `pkg-config --cflags --libs MagickWand`
  - 曰く
  - > Check if pkg-config is able to find the right ImageMagick include and libs

### 2. 環境変数 `CGO_CFLAGS_ALLOW` を設定

### ImageMagick-6 のポリシーを書き換える

脆弱性対策のために PDF の変換を許可しないポリシーが書いてあるので、コメントアウトする。

```xml
<!-- /etc/ImageMagick-6/policy.xml -->
<policymap>
  ...
  <policy domain="coder" rights="none" pattern="PDF" />
</policymap>
```

終わったら戻す。まぁローカルなら問題ないだろうけど。

---

## 参考

- [Go で PDF を画像に変換する - Qiita](https://qiita.com/toshikitsubouchi/items/51c3268185cdc976a52f)
