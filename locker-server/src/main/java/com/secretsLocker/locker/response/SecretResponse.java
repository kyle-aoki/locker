package com.secretsLocker.locker.response;

public class SecretResponse extends Response {
    public static class SecretCreated extends RepoResponse {
        public String code = "SR001";
        public String message = "Secret created.";
    }
    public static class SecretSent extends RepoResponse {
        public String code = "SR002";
        public String message = "Secret sent.";
        public String secret;
        public SecretSent(String secretToSend) {
            this.secret = secretToSend;
        }
    }
}
