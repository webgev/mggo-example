package controller

import (
    "github.com/webgev/mggo"
    "strconv"
)

func init() {
    mggo.AppendRight("Catalog.List", mggo.RRightGuest)
    mggo.AppendRight("Catalog.ListCategory", mggo.RRightGuest)
    mggo.AppendRight("Catalog.Read", mggo.RRightGuest)
    mggo.AppendRight("Catalog.BasketList", mggo.RRightGuest)
    mggo.AppendRight("Catalog.ReadCategory", mggo.RRightGuest)
    mggo.AppendRight("Catalog.Update", mggo.RRightEditor)
    mggo.AppendRight("Catalog.Create", mggo.RRightEditor)
    mggo.AppendRight("Catalog.Delete", mggo.RRightEditor)

    mggo.AppendViewRight("Catalog.Create", mggo.RRightEditor)

    mggo.InitCallback(func () {
        models := []interface{}{}
        models = append(models, (*Category)(nil))
        models = append(models, (*Basket)(nil))
        models = append(models, (*ProductCat)(nil))
        mggo.CreateTable(models)
    })   
}

type Catalog struct {
    ProductID  int
    CategoryID int
    UserID     int
    Product    Product
    View       CatalogView
}
type Product struct {
    ID          int
    Name        string
    Description string
    Price       float64
    Active      bool
}
type Category struct {
    ID          int
    Name        string
    Description string
}
type Basket struct {
    ID        int
    ProductID int
    Count     int
    UserID    int
}
type ProductCat struct {
    ProductID  int
    CategoryID int
}

func (c Catalog) List() (products []Product) {
    product := c.Product
    mggo.SQL().Model(product).Select(&products)
    return
}
func (c Catalog) Read() Product {
    product := Product{ID: c.ProductID}
    mggo.SQL().Select(&product)
    return product
}

func (c Catalog) ListCategory() (categories []Category) {
    mggo.SQL().Model(&Category{}).Select(&categories)
    return categories
}
func (c Catalog) ReadCategory() Category {
    category := Category{ID: c.CategoryID}
    mggo.SQL().Select(&category)
    return category
}
func (c Catalog) BasketList() (baskets []Basket) {
    basket := Basket{}
    if c.UserID != 0 {
        basket.UserID = c.UserID
    }
    mggo.SQL().Model(&basket).Select(&baskets)
    return
}

func (c Catalog) Update() int {
    if c.Product.ID == 0 {
        mggo.SQL().Insert(&c.Product)
    } else {
        mggo.SQL().Update(&c.Product)
    }
    return c.Product.ID
}

func (c Catalog) Create() int {
    return c.Update()
}

func (c Catalog) Delete() {
    if c.Product.ID != 0 {
        product := c.Product
        mggo.SQL().Delete(&product)
    }
}

func (p Product) List(catID int) (products []Product) {
    query := mggo.SQL().Model(&p)
    if catID > 0 {
        query.Join(`JOIN "product_cats" on "category_id" = ? and "product_id" = id`, catID)
    }
    query.Select(&products)
    return
}

func (c Category) Read() Category {
    if c.ID > 0 {
        mggo.SQL().Select(&c)
    }
    return c
}

type CatalogView struct{}

func (c CatalogView) Index(data *mggo.ViewData, path []string) {
    data.View = "catalog/catalog.html"
    data.Data["Title"] = "Catalog"
    data.Data["Categories"] = Catalog{}.ListCategory()
}

func (c CatalogView) Create(data *mggo.ViewData, path []string) {
    data.View = "catalog/create.html"
    data.Data["Title"] = "Catalog"
    data.Data["Model"] = Product{}
}

func (c CatalogView) Products(data *mggo.ViewData, path []string) {
    var catID int
    title := "Catalog"
    if len(path) > 2 {
        if i, err := strconv.Atoi(path[2]); err == nil {
            catID = i
            cat := Category{ID: catID}.Read()
            title = cat.Name
        }
    }
    data.View = "catalog/products.html"
    data.Data["Title"] = title
    data.Data["Products"] = Product{}.List(catID)
}
