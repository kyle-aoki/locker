package com.secretsLocker.locker.response;

import com.secretsLocker.locker.dto.RawSecret;

import java.util.List;

public class EnvironmentResponse {
    public static class EnvCreated extends Response {
        public String code = "ER001";
        public String message = "Environment created.";
    }
    public static class GetEnv extends Response {
        public String code = "ER002";
        public List<RawSecret> list;
        public GetEnv(List<RawSecret> list) {
            this.list = list;
        }
    }
}
