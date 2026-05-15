import Fluent

struct TestProduct: AsyncMigration {
    func prepare(on database: Database) async throws {
        try await Product(name: "test01", price: 11.11).save(on: database)
        try await Product(name: "test02", price: 2222.22).save(on: database)
    }

    func revert(on database: Database) async throws {
        try await Product.query(on: database).delete()
    }
}
