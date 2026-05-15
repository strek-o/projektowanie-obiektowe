import Vapor
import Fluent
import FluentSQLiteDriver

public func configure(_ app: Application) async throws {
    app.databases.use(.sqlite(.file("db.sqlite")), as: .sqlite)

    app.migrations.add(CreateProduct())
    app.migrations.add(TestProduct())

    try await app.autoMigrate()

    try routes(app)
}
