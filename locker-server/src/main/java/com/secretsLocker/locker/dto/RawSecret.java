package com.secretsLocker.locker.dto;

public class RawSecret {
    public String key;
    public String value;
    public RawSecret(String key, String value) {
        this.key = key;
        this.value = value;
    }
}
