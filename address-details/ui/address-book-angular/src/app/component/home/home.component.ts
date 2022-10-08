import { Router } from '@angular/router';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css'],
})
export class HomeComponent implements OnInit {
  isLoggedIn: any = true;
  constructor(private router: Router) {
    if (sessionStorage.getItem('token')) {
      console.log('logged in');
      this.isLoggedIn = false;
    } else {
      this.isLoggedIn = true;
      console.log('not logged in');
    }
  }

  ngOnInit(): void {}

  redirectLogin() {
    console.log('triggered');
    this.router.navigate(['/login']);
  }
  redirectRegister() {
    console.log('triggered');
    this.router.navigate(['/register']);
  }

  redirectLoginout() {
    this.router.navigate(['/logout']);
  }
}
