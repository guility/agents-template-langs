use async_trait::async_trait;
use ddd_rust_project::{
    BaseEntity, Controller, InMemoryRepository, Repository, RepositoryError, UseCase,
};

struct LookupUseCase {
    repository: InMemoryRepository<BaseEntity>,
}

#[async_trait]
impl UseCase for LookupUseCase {
    type Request = String;
    type Response = Option<BaseEntity>;
    type Error = RepositoryError;

    async fn execute(&self, request: String) -> Result<Self::Response, Self::Error> {
        self.repository.find_by_id(&request).await
    }
}

#[tokio::test]
async fn stored_account_is_returned_through_controller_response() {
    let repository = InMemoryRepository::<BaseEntity>::new();
    repository.save(BaseEntity::new("account-1")).await.unwrap();

    let response = Controller
        .handle_request("account-1".to_owned(), &LookupUseCase { repository })
        .await;

    assert!(response.success);
    assert_eq!(response.data.unwrap().unwrap().id, "account-1");
}
