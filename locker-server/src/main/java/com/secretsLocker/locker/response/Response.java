package com.secretsLocker.locker.response;

public class Response {
    public Boolean ok = true;
    public String code;

    public Response(String code) {
        this.code = code;
    }
}
