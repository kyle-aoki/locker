package com.secretsLocker.locker.response;

import java.util.List;

public class ListResponse extends Response {
    public List<String> list;

    public ListResponse(String code, List<String> list) {
        super(code);
        this.list = list;
    }
}
