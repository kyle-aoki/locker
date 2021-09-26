package com.secretsLocker.locker.response;

public class SecretResponse extends Response {
    public static class SecretCreated extends RepoResponse {
        public String code = "SR001";
        public String message = "Secret created.";
    }
    public static class SecretUpdated extends RepoResponse {
        public String code = "SR002";
        public String message = "Secret updated.";
    }
    public static class SecretSent extends RepoResponse {
        public String code = "SR003";
        public String message = "Secret sent.";
        public String secretValue;
        public SecretSent(String secretToSend) {
            this.secretValue = secretToSend;
        }
    }
}
