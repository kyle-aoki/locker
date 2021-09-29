package com.secretsLocker.locker.response;

import com.secretsLocker.locker.dto.KeyValue;

import java.util.List;

public class KeyValueResponse extends Response {
    public List<KeyValue> keyValueList;

    public KeyValueResponse(String code, List<KeyValue> keyValueList) {
        super(code);
        this.keyValueList = keyValueList;
    }
}
