package com.secretsLocker.locker.response;

public class EnvironmentResponse extends Response {
    public static class EnvCreated extends RepoResponse {
        public String code = "ER001";
        public String message = "Environment created.";
    }
}
