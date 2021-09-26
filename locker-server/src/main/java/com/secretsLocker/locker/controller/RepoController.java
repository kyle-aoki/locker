package com.secretsLocker.locker.controller;

import com.secretsLocker.locker.dto.CreateRepoDto;
import com.secretsLocker.locker.dto.ListRepoDto;
import com.secretsLocker.locker.entity.Repository;
import com.secretsLocker.locker.response.RepoResponse;
import com.secretsLocker.locker.service.RepoService;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/repo")
public class RepoController {

    Logger logger = LoggerFactory.getLogger(RepoController.class);

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

    @PostMapping("/list")
    public RepoResponse.RepoList list(@RequestBody ListRepoDto listRepoDto) {
        logger.info("Received request to list repositories");
        List<String> repos = repoService.listRepos(listRepoDto);
        return new RepoResponse.RepoList(repos);
    }

}
