package com.secretsLocker.locker.controller;

import com.secretsLocker.locker.dto.CreateSecretDto;
import com.secretsLocker.locker.dto.GetSecretDto;
import com.secretsLocker.locker.response.SecretResponse;
import com.secretsLocker.locker.service.SecretService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/secret")
public class SecretController {

    @Autowired
    SecretService secretService;

    @PostMapping("/create")
    public SecretResponse.SecretCreated createSecret(@RequestBody CreateSecretDto createSecretDto) {
        secretService.create(createSecretDto);
        return new SecretResponse.SecretCreated();
    }

    @PostMapping("/update")
    public SecretResponse.SecretUpdated update(@RequestBody CreateSecretDto createSecretDto) {
        secretService.update(createSecretDto);
        return new SecretResponse.SecretUpdated();
    }

    @PostMapping("/get")
    public SecretResponse.SecretSent getSecret(@RequestBody GetSecretDto getSecretDto) {
        String secretValue = secretService.get(getSecretDto);
        return new SecretResponse.SecretSent(secretValue);
    }
}
