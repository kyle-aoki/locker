package com.secretsLocker.locker.controller;

import com.secretsLocker.locker.dto.ListRepoDto;
import com.secretsLocker.locker.dto.copy.CopyRepoDto;
import com.secretsLocker.locker.dto.delete.DeleteRepoDto;
import com.secretsLocker.locker.dto.path.RepoPath;
import com.secretsLocker.locker.dto.UpdateRepoDto;
import com.secretsLocker.locker.response.ListResponse;
import com.secretsLocker.locker.response.MessageResponse;
import com.secretsLocker.locker.response.Response;
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
    public Response createRepo(
            @RequestHeader("username") String username,
            @RequestBody RepoPath repoPath
            ) {
        repoService.create(username, repoPath);
        return new MessageResponse("RC200", "Repository created.");
    }

    @PostMapping("/rename")
    public Response rename(
            @RequestHeader("username") String username,
            @RequestBody UpdateRepoDto updateRepoDto
    ) {
        repoService.rename(username, updateRepoDto);
        return new MessageResponse("RUN200", "Repository name updated.");
    }

    @PostMapping("/list")
    public Response list(@RequestBody ListRepoDto listRepoDto) {
        List<String> repos = repoService.listRepos(listRepoDto);
        return new ListResponse("RL200", repos);
    }

    @PostMapping("/delete")
    public Response delete(@RequestBody DeleteRepoDto deleteRepoDto) {
        repoService.delete(deleteRepoDto);
        return new MessageResponse("RP_DL_200", "Deleted repository " + deleteRepoDto.repoName + ".");
    }

    @PostMapping("/copy")
    public Response copy(@RequestBody CopyRepoDto copyRepoDto) {
        repoService.copy(copyRepoDto);
        return new MessageResponse("RP_CL_200", "Cloned repository " + copyRepoDto.repoName + ".");
    }
}
