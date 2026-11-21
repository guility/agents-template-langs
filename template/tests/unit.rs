use async_trait::async_trait;
use ddd_rust_project::{
    BaseEntity, Controller, Entity, InMemoryRepository, Repository, RepositoryError, ResponseDTO,
    UseCase,
};
use std::fmt::Debug;

#[derive(Clone, Debug)]
struct SampleEntity(BaseEntity);

impl Entity for SampleEntity {
    fn id(&self) -> &str {
        self.0.id()
    }
    fn created_at(&self) -> time::OffsetDateTime {
        self.0.created_at()
    }
    fn updated_at(&self) -> Option<time::OffsetDateTime> {
        self.0.updated_at()
    }
}

struct EchoUseCase;

#[async_trait]
impl UseCase for EchoUseCase {
    type Request = String;
    type Response = String;
    type Error = String;

    async fn execute(&self, request: String) -> Result<String, String> {
        if request.is_empty() {
            Err("empty request".into())
        } else {
            Ok(request)
        }
    }
}

#[tokio::test]
async fn entity_and_repository_cover_success_and_missing_paths() {
    let mut base = BaseEntity::new("id-1");
    assert_eq!(base.id(), "id-1");
    assert!(base.created_at().unix_timestamp() > 0);
    assert!(base.updated_at().is_none());
    base.update_timestamp();
    assert!(base.updated_at().is_some());

    let repository = InMemoryRepository::default();
    assert!(repository.find_by_id("missing").await.unwrap().is_none());
    repository.save(SampleEntity(base.clone())).await.unwrap();
    assert_eq!(
        repository.find_by_id("id-1").await.unwrap().unwrap().id(),
        "id-1"
    );
    assert_eq!(repository.find_all().await.unwrap().len(), 1);
    repository.delete("id-1").await.unwrap();
    assert!(repository.find_all().await.unwrap().is_empty());
}

#[tokio::test]
async fn controller_and_response_cover_success_and_error_paths() {
    let controller = Controller;
    let success = controller.handle_request("ok".into(), &EchoUseCase).await;
    assert_eq!(success.data.as_deref(), Some("ok"));
    assert!(success.success);

    let failure = controller.handle_request(String::new(), &EchoUseCase).await;
    assert_eq!(failure.error.as_deref(), Some("\"empty request\""));
    assert!(!failure.success);

    let response = ResponseDTO::<String>::success("data".into());
    assert_eq!(response.data.as_deref(), Some("data"));
    let response = ResponseDTO::<String>::error("failure");
    assert_eq!(response.error.as_deref(), Some("failure"));
}

#[test]
fn repository_error_formats_every_variant() {
    assert_eq!(
        RepositoryError::NotFound("x".into()).to_string(),
        "Not found: x"
    );
    assert_eq!(
        RepositoryError::Conflict("x".into()).to_string(),
        "Conflict: x"
    );
    assert_eq!(
        RepositoryError::Internal("x".into()).to_string(),
        "Internal error: x"
    );
}
