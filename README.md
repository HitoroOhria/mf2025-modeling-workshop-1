# MF2025 ワークショップ①『モデル駆動設計をやってみよう』

- [MF2025 ワークショップ①『モデル駆動設計をやってみよう』](https://umtp-japan.org/event-seminar/mf2025/76609)
- [モデル駆動設計をやってみよう Modeling Forum2025ワークショップ/Let’s Try Model-Driven Design](https://speakerdeck.com/haru860/lets-try-model-driven-design)

# シンプル版

## ユースケース

- 参加者が日程調整を作成する
    - 候補日を選択する
- 参加者が出欠を登録する
    - 候補日に対し、ユーザーが◯か✗かを入力する
    - 未記入の候補日があってはならない
- システムが最適な日付を確定する
    - 新しいユーザーが出欠を登録するたびに、システムによる最適な日付の確定が行われる
    - 最も◯が多い候補日を確定する
    - 最も丸が多い候補日が複数ある場合、複数の最適な日付を確定する

## モデル

参加者が出欠を登録するユースケース

```mermaid
graph LR
  subgraph group1[出欠記入]
	  user1("【佐藤さん】<br/>12月3日：◯<br/>12月7日：◯<br/>12月8日：✗")
	  user2("【鈴木さん】<br/>12月3日：◯<br/>12月7日：◯<br/>12月8日：◯")
	  user3("【高橋さん】<br/>12月3日：✗<br/>12月7日：◯<br/>12月8日：◯")
  end
```

システムが最適な日付を確定するユースケース

```mermaid
graph LR
  subgraph group2[出欠表]
	  proposedDate1("【出欠票】<br/>出欠1<br/>出欠2<br/>出欠3")
  end
  
  subgraph group3[出欠]
	  attendance1("【出欠1】<br/>12月3日<br/>佐藤さん<br/>◯")
	  attendance2("【出欠2】<br/>12月7日<br/>佐藤さん<br/>◯")
	  attendance3("【出欠3】<br/>12月8日<br/>佐藤さん<br/>✗")
  end
  
  subgraph group4[候補日グループ - 仮]
    dateGroup1("【12月3日】<br/>候補日グループアイテム1")
    dateGroup2("【12月7日】<br/>候補日グループアイテム2")
    dateGroup3("【12月8日】<br/>候補日グループアイテム3")
  end

  subgraph group5[候補日グループアイテム - 仮]
    dateGroupItem1("【候補日グループアイテム1】<br/>佐藤さん<br/>◯")
    dateGroupItem2("【候補日グループアイテム2】<br/>佐藤さん<br/>◯")
    dateGroupItem3("【候補日グループアイテム3】<br/>佐藤さん<br/>◯")
  end

  group2-->group3
  group2-. 中間モデル .->group4
  group4-->group5
```

## 振り返り

- 「候補日グループ - 仮」の中間モデルを用意するか？
  - 「週ごとに出欠をまとめ、カウントする処理」は専用モデルがあった方がコードを理解しやすいと考えた
  - しかし、「候補日グループ - 仮」はドメイン上のモデルではなく、あくまでもプログラミングを助けるための中間モデルでしかない
