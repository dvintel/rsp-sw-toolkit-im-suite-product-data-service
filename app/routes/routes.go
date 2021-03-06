/* Apache v2 license
*  Copyright (C) <2019> Intel Corporation
*
*  SPDX-License-Identifier: Apache-2.0
 */
package routes

import (
	"database/sql"

	"github.com/gorilla/mux"

	"github.com/intel/rsp-sw-toolkit-im-suite-product-data-service/app/routes/handlers"
	"github.com/intel/rsp-sw-toolkit-im-suite-product-data-service/pkg/middlewares"
	"github.com/intel/rsp-sw-toolkit-im-suite-product-data-service/pkg/web"
)

// Route struct holds attributes to declare routes
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc web.Handler
}

// NewRouter creates the routes for GET and POST
func NewRouter(db *sql.DB, size int) *mux.Router {

	mapp := handlers.Mapping{MasterDB: db, Size: size}

	var routes = []Route{
		// swagger:operation GET / default Healthcheck
		//
		// Healthcheck Endpoint
		//
		// Endpoint that is used to determine if the application is ready to take web requests
		//
		// ---
		// consumes:
		// - application/json
		//
		// produces:
		// - application/json
		//
		//
		// schemes:
		// - http
		//
		// responses:
		//   '200':
		//     description: OK
		//
		{
			"Index",
			"GET",
			"/",
			mapp.Index,
		},
		// swagger:route POST /skus skus postSkus
		//
		// Loads SKU Data
		//
		// This API call is used to upload a list of SKU items into Responsive Retail.<br>
		// The SKU data includes:
		//
		// <blockquote>• <b>SKU</b>: The SKU number, a unique identifier for the SKU.</blockquote>
		//
		// <blockquote>• <b>Product List</b>: A list (array) of UPCs that are included in the SKU.</blockquote>
		//
		// Expected formatting of JSON input (as an example):<br><br>
		//
		//```json
		// {
		//"data":[{
		//   "sku" : "MS122-32",
		//   "productList" : [
		//     { "productId": "00888446671444", "metadata": {"color":"blue"} },
		//     { "productId": "889319762751", "metadata": {"size":"small"} }
		//    ]
		// 	 },
		// 	{
		//    "sku" : "MS122-34",
		//    "productList" : [
		//      {"productId": "90388987132758", "metadata": {"name":"pants"} }
		//    ]
		//  }]
		// }
		//```
		// <br>
		// Each SKU item is treated individually; it succeeds or fails independent of the other SKUs.
		// Check the returned results to determine the success or failure of each SKU.
		//
		//     Consumes:
		//     - application/json
		//
		//     Produces:
		//     - application/json
		//
		//     Schemes: http
		//
		//
		//     Responses:
		//       201: Created
		//       400: schemaValidation
		//       500: internalError
		//
		{
			"PostSkuMapping",
			"POST",
			"/skus",
			mapp.PostSkuMapping,
		},
		// swagger:route GET /skus skus getSkus
		//
		// Retrieves SKU Data
		//
		// This API call is used to retrieve a list of SKU items.
		//
		// <blockquote>• <b>Search by sku</b>: To search by sku, you would use the filter query parameter like so: /sku?$filter=(sku eq 'MS122-32')</blockquote>
		//
		// <blockquote>• <b>Search by name</b>: To search by name, you would use the filter query parameter like so: /location?$filter=(name eq 'mens khaki slacks')</blockquote>
		//
		//
		// `/skus?$top=10&$select=sku` - Useful for paging data. Grab the top 10 records and only pull back the sku field
		//
		// `/skus?$count` - Tell me how many records are in the database
		//
		// `/skus?$filter=(sku eq '12345678') and (productList.metadata.color eq 'red')` - This filters on particular sku and UPCs that are classified as "Red"
		//
		// `/skus?$orderby=sku desc` - Give me back all skus in descending order by sku
		//
		// `/skus?$filter=startswith(sku,'m')` - Give me all skus that begin with the letter 'm'
		//
		// `/skus?$count&$filter=(sku eq '12345678')` - Give me the count of items with the SKU `12345678``
		//
		// `/skus?$inlinecount=allpages&$filter=(sku eq '12345678')` - Give me all items with the SKU `12345678` and include how many there are
		//
		//
		//
		// Example Result:<br><br>
		//```json
		// {
		//     "results": [
		//         {
		//             "sku": "12345679",
		//             "productList": [
		//                 {
		//                     "metadata": {
		//                         "color": "blue",
		//                         "size": "XS"
		//                     },
		//                     "productId": "123456789783"
		//                 },
		//                 {
		//                     "metadata": {
		//                         "color": "red",
		//                         "size": "M"
		//                     },
		//                     "productId": "123456789784"
		//                 }
		//             ]
		//         }
		//     ]
		// }
		//```
		//
		//     Consumes:
		//     - application/json
		//
		//     Produces:
		//     - application/json
		//
		//     Schemes: http
		//
		//     Responses:
		//       200: body:resultsResponse
		//       400: schemaValidation
		//       500: internalError
		//
		{
			"GetSkuMapping",
			"GET",
			"/skus",
			mapp.GetSkuMapping,
		},
		// swagger:route GET /productid/{productid} productid productids
		//
		// Retrieves SKU Data
		//
		// This API call is used to get the metadata for a upc.<br><br>
		//
		// Example query:
		//
		// <blockquote>/productid/12345678978345</blockquote> <br><br>
		//
		//
		// Example Result: <br><br>
		//```json
		// {
		//   "metadata": {
		// 				"color": "blue",
		// 				 "size": "XS"
		// 			  },
		//   "productid": "12345678978345"
		// }
		//```
		//
		//     Consumes:
		//     - application/json
		//
		//     Produces:
		//     - application/json
		//
		//     Schemes: http
		//
		//     Responses:
		//       200: body:resultsResponse
		//       404: NotFound
		//       400: schemaValidation
		//       500: internalError
		//
		{
			"GetProductID",
			"GET",
			"/productid/{productId}",
			mapp.GetProductID,
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		handler := route.HandlerFunc
		handler = middlewares.Recover(handler)
		handler = middlewares.Logger(handler)
		handler = middlewares.BodyLimiter(handler)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
