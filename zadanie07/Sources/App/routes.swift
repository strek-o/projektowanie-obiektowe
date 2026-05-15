import Vapor

func routes(_ app: Application) throws {
    app.get { req async in
        "GOTO: /products"
    }

    try app.register(collection: ProductController())
}
