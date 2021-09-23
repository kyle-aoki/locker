package com.secretsLocker.locker.controller;

import com.secretsLocker.locker.dto.CreateRepoDto;
import com.secretsLocker.locker.response.RepoResponse;
import com.secretsLocker.locker.service.RepoService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/repo")
public class RepoController {

    @Autowired
    RepoService repoService;

    @PostMapping("/create")
    public RepoResponse.RepoCreated createRepo(
            @RequestHeader("username") String username,
            @RequestBody CreateRepoDto createRepoDto
    ) {
        repoService.create(username, createRepoDto);
        return new RepoResponse.RepoCreated();
    }

}
