package com.secretsLocker.locker.response;

import java.util.List;

public class SecretResponse {
    public static class SecretCreated extends Response {
        public String code = "SR001";
        public String message = "Secret created.";
    }
    public static class SecretUpdated extends Response {
        public String code = "SR002";
        public String message = "Secret updated.";
    }
    public static class SecretSent extends Response {
        public String code = "SR003";
        public String message = "Secret sent.";
        public String secretValue;
        public SecretSent(String secretToSend) {
            this.secretValue = secretToSend;
        }
    }
    public static class SecretNameList extends Response {
        public String code = "SR004";
        public List<String> list;
        public SecretNameList(List<String> secretNameList) {
            this.list = secretNameList;
        }
    }
    public static class SecretValueList extends Response {
        public String code = "SR005";
        public List<String> list;
        public SecretValueList(List<String> secretNameList) {
            this.list = secretNameList;
        }
    }
}
