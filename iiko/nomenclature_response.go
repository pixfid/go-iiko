/*
 * MIT License
 *
 * Copyright (c) 2020 Semchenko Aleksandr
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package iiko

type NomenclatureResponse struct {
	Groups            []Group           `json:"groups"`
	ProductCategories []ProductCategory `json:"productCategories"`
	Products          []Product         `json:"products"`
	Revision          *int64            `json:"revision,omitempty"`
	UploadDate        *string           `json:"uploadDate,omitempty"`
}

type Group struct {
	AdditionalInfo   interface{}   `json:"additionalInfo"`
	Code             interface{}   `json:"code"`
	Description      interface{}   `json:"description"`
	ID               *string       `json:"id,omitempty"`
	IsDeleted        *bool         `json:"isDeleted,omitempty"`
	Name             *string       `json:"name,omitempty"`
	SEODescription   interface{}   `json:"seoDescription"`
	SEOKeywords      interface{}   `json:"seoKeywords"`
	SEOText          interface{}   `json:"seoText"`
	SEOTitle         interface{}   `json:"seoTitle"`
	Tags             interface{}   `json:"tags"`
	Images           []interface{} `json:"images"`
	IsIncludedInMenu *bool         `json:"isIncludedInMenu,omitempty"`
	Order            *int64        `json:"order,omitempty"`
	ParentGroup      *string       `json:"parentGroup"`
}

type ProductCategory struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type Product struct {
	AdditionalInfo         interface{}          `json:"additionalInfo"`
	Code                   *string              `json:"code,omitempty"`
	Description            *string              `json:"description,omitempty"`
	ID                     *string              `json:"id,omitempty"`
	IsDeleted              *bool                `json:"isDeleted,omitempty"`
	Name                   *string              `json:"name,omitempty"`
	SEODescription         interface{}          `json:"seoDescription"`
	SEOKeywords            interface{}          `json:"seoKeywords"`
	SEOText                interface{}          `json:"seoText"`
	SEOTitle               interface{}          `json:"seoTitle"`
	Tags                   interface{}          `json:"tags"`
	CarbohydrateAmount     *float64             `json:"carbohydrateAmount,omitempty"`
	CarbohydrateFullAmount *float64             `json:"carbohydrateFullAmount,omitempty"`
	DifferentPricesOn      []DifferentPricesOn  `json:"differentPricesOn"`
	DoNotPrintInCheque     *bool                `json:"doNotPrintInCheque,omitempty"`
	EnergyAmount           *float64             `json:"energyAmount,omitempty"`
	EnergyFullAmount       *float64             `json:"energyFullAmount,omitempty"`
	FatAmount              *float64             `json:"fatAmount,omitempty"`
	FatFullAmount          *float64             `json:"fatFullAmount,omitempty"`
	FiberAmount            *float64             `json:"fiberAmount,omitempty"`
	FiberFullAmount        *float64             `json:"fiberFullAmount,omitempty"`
	GroupID                *string              `json:"groupId"`
	GroupModifiers         []GroupModifier      `json:"groupModifiers"`
	MeasureUnit            *string              `json:"measureUnit,omitempty"`
	Modifiers              []Modifier           `json:"modifiers"`
	Price                  *int64               `json:"price,omitempty"`
	ProductCategoryID      *string              `json:"productCategoryId,omitempty"`
	ProhibitedToSaleOn     []ProhibitedToSaleOn `json:"prohibitedToSaleOn"`
	Type                   *string              `json:"type,omitempty"`
	UseBalanceForSell      *bool                `json:"useBalanceForSell,omitempty"`
	Weight                 *float64             `json:"weight,omitempty"`
	Images                 []Image              `json:"images"`
	IsIncludedInMenu       *bool                `json:"isIncludedInMenu,omitempty"`
	Order                  *int64               `json:"order,omitempty"`
	ParentGroup            *string              `json:"parentGroup"`
	WarningType            *int64               `json:"warningType,omitempty"`
}

type DifferentPricesOn struct {
	Price         *int64      `json:"price,omitempty"`
	PriceCategory interface{} `json:"priceCategory"`
	TerminalID    *string     `json:"terminalId,omitempty"`
}

type GroupModifier struct {
	MaxAmount                            *int64     `json:"maxAmount,omitempty"`
	MinAmount                            *int64     `json:"minAmount,omitempty"`
	ModifierID                           *string    `json:"modifierId,omitempty"`
	Required                             *bool      `json:"required,omitempty"`
	ChildModifiers                       []Modifier `json:"childModifiers"`
	ChildModifiersHaveMinMaxRestrictions *bool      `json:"childModifiersHaveMinMaxRestrictions,omitempty"`
}

type Modifier struct {
	MaxAmount           *int64  `json:"maxAmount,omitempty"`
	MinAmount           *int64  `json:"minAmount,omitempty"`
	ModifierID          *string `json:"modifierId,omitempty"`
	Required            *bool   `json:"required,omitempty"`
	DefaultAmount       *int64  `json:"defaultAmount,omitempty"`
	HideIfDefaultAmount *bool   `json:"hideIfDefaultAmount,omitempty"`
}

type Image struct {
	ImageID    *string `json:"imageId,omitempty"`
	ImageURL   *string `json:"imageUrl,omitempty"`
	UploadDate *string `json:"uploadDate,omitempty"`
}

type ProhibitedToSaleOn struct {
	TerminalID *string `json:"terminalId,omitempty"`
}
