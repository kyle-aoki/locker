package com.secretsLocker.locker.response;

public class StringResponse extends Response {
    public String str;
    public StringResponse(String code, String str) {
        super(code);
        this.str = str;
    }
}
