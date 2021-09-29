package com.secretsLocker.locker.controller;

import com.secretsLocker.locker.dto.*;
import com.secretsLocker.locker.response.ListResponse;
import com.secretsLocker.locker.response.MessageResponse;
import com.secretsLocker.locker.response.Response;
import com.secretsLocker.locker.response.StringResponse;
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
    public Response createSecret(@RequestBody CreateSecretDto createSecretDto) {
        secretService.create(createSecretDto);
        return new MessageResponse("SC200", "Secret created.");
    }

    @PostMapping("/update")
    public Response update(@RequestBody CreateSecretDto createSecretDto) {
        secretService.update(createSecretDto);
        return new MessageResponse("SU200", "Secret updated.");
    }

    @PostMapping("/rename")
    public Response rename(@RequestBody RenameSecretDto renameSecretDto) {
        secretService.rename(renameSecretDto);
        return new MessageResponse("SR200", "Secret renamed.");
    }

    @PostMapping("/get")
    public Response get(@RequestBody GetSecretDto getSecretDto) {
        String secretValue = secretService.get(getSecretDto);
        return new StringResponse("SG200", secretValue);
    }

    @PostMapping("/get-all")
    public Response getAll(@RequestBody GetAllSecretsDto getAllSecretsDto) {
        List<String> secretValues = secretService.getAll(getAllSecretsDto);
        return new ListResponse("SGA200", secretValues);
    }

    @PostMapping("/list")
    public Response list(@RequestBody ListSecretsDto listSecretsDto) {
        List<String> secretNameList = secretService.list(listSecretsDto);
        return new ListResponse("SL200", secretNameList);
    }
}
