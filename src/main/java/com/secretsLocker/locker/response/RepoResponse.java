package com.secretsLocker.locker.response;

public class RepoResponse extends Response {
    public static class RepoCreated extends RepoResponse {
        public String code = "RR001";
        public String message = "Repository created.";
    }
}
