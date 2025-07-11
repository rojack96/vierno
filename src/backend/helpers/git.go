package helpers

import (
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

type Git struct {
	Repo        string  `json:"repo" yaml:"repo"`
	Branch      string  `json:"branch" yaml:"branch"`
	RemoteName  *string `json:"remoteName" yaml:"remoteName"`
	AccessToken *string `json:"accessToken" yaml:"accessToken"`
	Folder      string  `json:"folder" yaml:"folder"`
	SshAuth     *struct {
		Agent          *bool   `json:"agent" yaml:"agent"`
		PrivateKeyPath *string `json:"privateKeyPath" yaml:"privateKeyPath"`
		Passwd         *string `json:"passwd" yaml:"passwd"`
	} `json:"sshAuth" yaml:"sshAuth"`
}

type GitOperations interface {
	Clone()
	Pull()
	Push()
	Commit()
	Checkout(branch string)
}

// Clone clones the git repository to the specified folder.
func (g *Git) Clone() {
	var (
		auth transport.AuthMethod
		err  error
	)
	fmt.Println("Clonazione in corso...") // todo traduzione
	if auth, err = g.auth(); err != nil {
		log.Fatalf("Errore durante l'autenticazione: %v", err) // todo traduzione
	}

	_, err = git.PlainClone(g.Folder, false, &git.CloneOptions{
		URL:        g.Repo,
		Progress:   nil,
		Auth:       auth,
		RemoteName: g.remoteName(),
	})

	if err != nil {
		log.Fatalf("Errore durante la clonazione: %v", err) // todo traduzione
	}

	fmt.Println("✅ Clonazione completata.") // todo traduzione
}

/*
// Pull pulls the latest changes from the remote repository.
func (g *Git) Pull() {
	var (
		auth transport.AuthMethod
		err  error
	)
	if auth, err = g.auth(); err != nil {
		log.Fatalf("Errore durante l'autenticazione: %v", err) // todo traduzione
	}
	fmt.Println("Pull in corso...")
	repo, err := git.PlainOpen(g.Folder)
	if err != nil {
		log.Fatalf("Errore durante l'apertura del repo: %v", err)
	}

	w, err := repo.Worktree()
	if err != nil {
		log.Fatalf("Errore durante l'ottenimento del worktree: %v", err)
	}

	err = w.Pull(&git.PullOptions{
		RemoteName: g.remoteName(),
		Auth:       auth,
	})
	if err != nil && err != git.NoErrAlreadyUpToDate {
		log.Fatalf("Errore durante il pull: %v", err)
	}

	fmt.Println("✅ Pull completato.")
}

func (g *Git) Push() {
	var (
		auth transport.AuthMethod
		err  error
	)
	if auth, err = g.auth(); err != nil {
		log.Fatalf("Errore durante l'autenticazione: %v", err) // todo
	}
	fmt.Println("Push in corso...")
	repo, err := git.PlainOpen(g.Folder)
	if err != nil {
		log.Fatalf("Errore durante l'apertura del repo: %v", err)
	}
	w, err := repo.Worktree()
	if err != nil {
		log.Fatalf("Errore durante l'ottenimento del worktree: %v", err)
	}
	err = w.Push(&git.PushOptions{
		RemoteName: g.remoteName(),
		Auth:       auth,
	})
	if err != nil {
		log.Fatalf("Errore durante il push: %v", err)
	}
	fmt.Println("✅ Push completato.")
}

// Checkout checks out the specified branch in the git repository.
func (g *Git) Checkout(branch string) {
	repo, err := git.PlainOpen(g.Folder)
	if err != nil {
		log.Fatalf("Errore durante l'apertura del repo: %v", err)
	}
	w, err := repo.Worktree()
	if err != nil {
		log.Fatalf("Errore durante l'ottenimento del worktree: %v", err)
	}

	branchRefName := plumbing.NewBranchReferenceName(branch)
	err = w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName(branchRefName),
		Force:  true,
	})
	if err != nil {
		log.Fatalf("Errore durante il checkout del branch %s: %v", g.Branch, err)
	}
	fmt.Printf("✅ Checkout al branch %s completato.\n", g.Branch)
}
*/
// remoteName returns the name of the remote repository.
func (g *Git) remoteName() string {
	result := "origin"
	if g.RemoteName != nil && *g.RemoteName != "" {
		result = *g.RemoteName
	}
	return result
}

// auth returns the authentication method for the git operations.
func (g *Git) auth() (transport.AuthMethod, error) {
	if g.AccessToken != nil && *g.AccessToken != "" {
		return &http.BasicAuth{Username: "vierno-config-server", Password: *g.AccessToken}, nil
	}

	if g.SshAuth == nil {
		return nil, fmt.Errorf("SSH authentication configuration is missing")
	}

	if g.SshAuth.Agent != nil && *g.SshAuth.Agent {
		authMethod, err := ssh.NewSSHAgentAuth("git")
		if err != nil {
			fmt.Printf("Failed to create SSH agent auth: %s\n", err.Error())
			return nil, err
		}
		return authMethod, nil
	}

	if g.SshAuth.PrivateKeyPath != nil {
		_, err := os.Stat(*g.SshAuth.PrivateKeyPath)
		if err != nil {
			fmt.Printf("read file %s failed %s\n", *g.SshAuth.PrivateKeyPath, err.Error())
			return nil, fmt.Errorf("private key file not found: %s", *g.SshAuth.PrivateKeyPath)
		}

		passwd := ""
		if g.SshAuth.Passwd != nil {
			passwd = *g.SshAuth.Passwd
		}

		publicKeys, _ := ssh.NewPublicKeysFromFile("git", *g.SshAuth.PrivateKeyPath, passwd)
		return publicKeys, nil
	}

	return nil, fmt.Errorf("no valid authentication method provided")
}
