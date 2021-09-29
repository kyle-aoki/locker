package com.secretsLocker.locker.response;

public class MessageResponse extends Response {
    public String message;

    public MessageResponse(String code, String message) {
        super(code);
        this.message = message;
    }
}
