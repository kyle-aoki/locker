package com.secretsLocker.locker.dto.diff;

import java.util.List;

public class DiffListDto {
    public List<MissingInOtherEnvs> otherMissing;
    public List<String> selfMissing;

    public DiffListDto(List<MissingInOtherEnvs> otherMissing, List<String> selfMissing) {
        this.otherMissing = otherMissing;
        this.selfMissing = selfMissing;
    }
}

