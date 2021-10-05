package com.secretsLocker.locker.controller;

import com.secretsLocker.locker.dto.*;
import com.secretsLocker.locker.dto.path.RepoEnvPath;
import com.secretsLocker.locker.dto.path.RepoEnvSecretPath;
import com.secretsLocker.locker.dto.path.RepoEnvSecretValuePath;
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
    public Response createSecret(@RequestBody RepoEnvSecretValuePath repoEnvSecretValuePath) {
        secretService.create(repoEnvSecretValuePath);
        return new MessageResponse("SC200", "Secret created.");
    }

    @PostMapping("/update")
    public Response update(@RequestBody RepoEnvSecretValuePath repoEnvSecretValuePath) {
        secretService.update(repoEnvSecretValuePath);
        return new MessageResponse("SU200", "Secret updated.");
    }

    @PostMapping("/rename")
    public Response rename(@RequestBody RenameSecretDto renameSecretDto) {
        secretService.rename(renameSecretDto);
        return new MessageResponse("SR200", "Secret renamed.");
    }

    @PostMapping("/get")
    public Response get(@RequestBody RepoEnvSecretPath repoEnvSecretPath) {
        String secretValue = secretService.get(repoEnvSecretPath);
        return new StringResponse("SG200", secretValue);
    }

    @PostMapping("/list")
    public Response list(@RequestBody RepoEnvPath repoEnvPath) {
        List<String> secretNameList = secretService.list(repoEnvPath);
        return new ListResponse("SL200", secretNameList);
    }

    @PostMapping("/delete")
    public Response delete(@RequestBody RepoEnvSecretPath repoEnvSecretPath) {
        secretService.delete(repoEnvSecretPath);
        return new MessageResponse("SD200", "Secret deleted.");
    }
}
