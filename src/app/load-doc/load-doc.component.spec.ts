import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LoadDocComponent } from './load-doc.component';

describe('LoadDocComponent', () => {
  let component: LoadDocComponent;
  let fixture: ComponentFixture<LoadDocComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [LoadDocComponent]
    });
    fixture = TestBed.createComponent(LoadDocComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
