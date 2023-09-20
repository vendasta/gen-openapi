import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LoadJsonComponent } from './load-json.component';

describe('LoadJsonComponent', () => {
  let component: LoadJsonComponent;
  let fixture: ComponentFixture<LoadJsonComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [LoadJsonComponent]
    });
    fixture = TestBed.createComponent(LoadJsonComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
