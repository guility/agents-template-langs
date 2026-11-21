use ddd_rust_project::{BaseEntity, InMemoryRepository, Repository};

#[tokio::test]
async fn repository_persists_domain_entities_across_calls() {
    let repository = InMemoryRepository::<BaseEntity>::new();
    let entity = BaseEntity::new("account-1");

    repository.save(entity.clone()).await.unwrap();

    let found = repository.find_by_id("account-1").await.unwrap().unwrap();
    assert_eq!(found.id, "account-1");
    assert_eq!(repository.find_all().await.unwrap().len(), 1);
}
