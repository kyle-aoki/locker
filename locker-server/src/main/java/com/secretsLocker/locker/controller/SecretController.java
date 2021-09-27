package com.secretsLocker.locker.controller;

import com.secretsLocker.locker.dto.CreateSecretDto;
import com.secretsLocker.locker.dto.GetAllSecretsDto;
import com.secretsLocker.locker.dto.GetSecretDto;
import com.secretsLocker.locker.dto.ListSecretsDto;
import com.secretsLocker.locker.response.SecretResponse;
import com.secretsLocker.locker.service.SecretService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.List;

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
    public SecretResponse.SecretSent get(@RequestBody GetSecretDto getSecretDto) {
        String secretValue = secretService.get(getSecretDto);
        return new SecretResponse.SecretSent(secretValue);
    }

    @PostMapping("/get-all")
    public SecretResponse.SecretValueList getAll(@RequestBody GetAllSecretsDto getAllSecretsDto) {
        List<String> secretValues = secretService.getAll(getAllSecretsDto);
        return new SecretResponse.SecretValueList(secretValues);
    }

    @PostMapping("/list")
    public SecretResponse.SecretNameList list(@RequestBody ListSecretsDto listSecretsDto) {
        List<String> secretNameList = secretService.list(listSecretsDto);
        return new SecretResponse.SecretNameList(secretNameList);
    }
}
