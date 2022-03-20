# 👺 Gilded Rose Refactoring Kata!

Hi and welcome to team **Gilded Rose**. As you know, we are a small inn with a prime location in a
prominent city ran by a friendly innkeeper named **Allison**. We also buy and sell only the finest goods.
Unfortunately, our goods are constantly degrading in `quality` as they approach their sell by date. 

We have a system in place that updates our `inventory` for us. It was developed by a no-nonsense type named
**Leeroy**, who has moved on to new adventures. Your task is to add the new feature to our system so that
we can begin selling a new category of `items`.

[Original Repository](https://github.com/emilybache/GildedRose-Refactoring-Kata)

## Functionalities
- [x] The value of an `Item` decrements by `1` each day
- [x] Once the `SellIn` has expired `Item` `Quality` decreases twice as fast
- [x] The `Quality` of an `Item` is never negative
- [x] The `Aged Brie` increments its `Quality` the older it gets
- [x] The `Quality` of an `Item` never surpasses `50`
- [x] The `Item` `Sulfuras, Hand of Ragnaros` its a legendary article, therefore its values are never modified
- [x] `Backstage passes` increment its `Quality` as `SellIn` approximates
  - [x] If there are `10` days left or less for the concert `Quality` increments by `2`
  - [x] If there are `5` days left or less for the concert `Quality` increments by `3`
  - [x] If the sell in time has expired `Quality` drops to `0`
- [x] `Conjured` `Items` degrade twice as fast than the normal `Items` 

## 🏃 Running the project

```bash
go run cmd/update/main.go
```

## 🤖 Running the tests

```bash
go test -v ./...
```

## License
The MIT License (MIT) - see [`license`](https://github.com/ksrof/gilded-rose-kata/blob/main/LICENSE) for more details.
