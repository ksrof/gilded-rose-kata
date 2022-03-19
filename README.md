# ▶️ How to run this project

## First run the project using the following command

```bash
go run cmd/update/main.go
```

**Expected output:**

```
------ Complex ------
Name: Aged Brie, SellIn: 1, Quality: 1
Name: Sulfuras, Hand of Ragnaros, SellIn: 0, Quality: 80
Name: Backstage passes to a TAFKAL80ETC concert, SellIn: 14, Quality: 21

------ Conjured ------
Name: Conjured Mana Cake, SellIn: 2, Quality: 4

------ Normal ------
Name: +5 Dexterity Vest, SellIn: 9, Quality: 19
Name: Elixir of the Mongoose, SellIn: 4, Quality: 6

------ Complex ------
Name: Aged Brie, SellIn: 0, Quality: 3
Name: Sulfuras, Hand of Ragnaros, SellIn: 0, Quality: 80
Name: Backstage passes to a TAFKAL80ETC concert, SellIn: 13, Quality: 22

------ Conjured ------
Name: Conjured Mana Cake, SellIn: 1, Quality: 2

------ Normal ------
Name: +5 Dexterity Vest, SellIn: 8, Quality: 18
Name: Elixir of the Mongoose, SellIn: 3, Quality: 5

------ Complex ------
Name: Aged Brie, SellIn: -1, Quality: 5
Name: Sulfuras, Hand of Ragnaros, SellIn: 0, Quality: 80
Name: Backstage passes to a TAFKAL80ETC concert, SellIn: 12, Quality: 23

------ Conjured ------
Name: Conjured Mana Cake, SellIn: 0, Quality: 0

------ Normal ------
Name: +5 Dexterity Vest, SellIn: 7, Quality: 17
Name: Elixir of the Mongoose, SellIn: 2, Quality: 4

------ Complex ------
Name: Aged Brie, SellIn: -2, Quality: 7
Name: Sulfuras, Hand of Ragnaros, SellIn: 0, Quality: 80
Name: Backstage passes to a TAFKAL80ETC concert, SellIn: 11, Quality: 24

------ Conjured ------
Name: Conjured Mana Cake, SellIn: -1, Quality: 0

------ Normal ------
Name: +5 Dexterity Vest, SellIn: 6, Quality: 16
Name: Elixir of the Mongoose, SellIn: 1, Quality: 3

------ Complex ------
Name: Aged Brie, SellIn: -3, Quality: 9
Name: Sulfuras, Hand of Ragnaros, SellIn: 0, Quality: 80
Name: Backstage passes to a TAFKAL80ETC concert, SellIn: 10, Quality: 26

------ Conjured ------
Name: Conjured Mana Cake, SellIn: -2, Quality: 0

------ Normal ------
Name: +5 Dexterity Vest, SellIn: 5, Quality: 15
Name: Elixir of the Mongoose, SellIn: 0, Quality: 0
```

## Then you can proceed to run the tests

```bash
go test -v ./...
```